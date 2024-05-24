import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma.service';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
	constructor(
		private readonly prismaService: PrismaService,
		private readonly jwtService: JwtService,
	) { }

	async login(data: { email: string; password: string }): Promise<string | Error> {
		const hashedPassword = await bcrypt.hash(data.password, await bcrypt.genSalt());
		if (this.prismaService.user.findUnique({
			where: {
				email: data.email,
				password: hashedPassword
			}
		})) {
			return this.generateAccessToken(data.email);
		}
		return new Error('Invalid credentials');
	}

	generateAccessToken(email: string): string {
		return this.jwtService.sign({ email });
	}
}
