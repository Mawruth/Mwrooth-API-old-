import { Injectable } from '@nestjs/common';
import { Prisma, User } from '@prisma/client';
import { PrismaService } from 'src/prisma.service';
import * as bcrypt from 'bcrypt';
import { randomInt } from 'crypto';
import { generateOtpVerificationMessage } from 'src/utils/utils';
import { MailerService } from '@nestjs-modules/mailer';
@Injectable()
export class UserService {
  constructor(
    private readonly prismaService: PrismaService,
    private readonly mailService: MailerService,
  ) { }
  
  async create(data: Prisma.UserCreateInput): Promise<User> {
    data.password = await bcrypt.hash(data.password, await bcrypt.genSalt());
    const otp = this.generateOtp() + "|5|" + Date.now().toString();
    data.otp = otp;
    const verificationMsg = generateOtpVerificationMessage(otp);
    await this.sendOtp(data.email, verificationMsg);
    return this.prismaService.user.create({ data });
  }

  async findAll(): Promise<User[]> {
    return this.prismaService.user.findMany({});
  }

  async findOneById(id: number): Promise<User> {
    return this.prismaService.user.findUnique({ where: { id } });
  }

  async findOneByEmail(email: string): Promise<User> {
    return this.prismaService.user.findUnique({ where: { email } });
  }

  async update(id: number, data: Prisma.UserUpdateInput): Promise<User> {
    return this.prismaService.user.update({ where: { id }, data });
  }

  async delete(id: number): Promise<User> {
    return this.prismaService.user.delete({ where: { id } });
  }

  async resendOtp(email: string): Promise<string | Error> {
    const otp = this.generateOtp() + "|5|" + Date.now().toString();
    const verificationMsg = generateOtpVerificationMessage(otp);
    this.sendOtp(email, verificationMsg).catch(err => {
      return err;
    });
    return "OTP resent";
  }

  async verifyOtp(email: string, otp: string): Promise<string | Error> {
    const user = await this.prismaService.user.findUnique({ where: { email } });
    if (!user) {
      return new Error("User not found");
    }
    const parsedOtp = this.parseOtp(user.otp);
    if (parsedOtp[0] !== otp) {
      return new Error("Invalid OTP");
    }
    if (Date.now() - parseInt(parsedOtp[2]) > parseInt(parsedOtp[1]) * 60 * 1000) {
      return new Error("OTP expired");
    }
    return new Error("OTP verified");
  }

  generateOtp(): string {
    const code = randomInt(0, 10000);
    return code.toString();
  }

  parseOtp(otp: string): string[] {
    return otp.split("|");
  }

  async sendOtp(email: string, msg: string) {
    await this.mailService.sendMail({
      from: process.env.SENDER_EMAIL,
      to: email,
      subject: 'OTP Verification',
      text: msg,
    })
  }
}
