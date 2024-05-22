import { Injectable } from '@nestjs/common';
import { Prisma, User } from '@prisma/client';
import { PrismaService } from 'src/prisma.service';

@Injectable()
export class UserService {
  constructor(
    private readonly prismaService: PrismaService
  ) { }
  create(data: Prisma.UserCreateInput): Promise<User> {
    return this.prismaService.user.create({ data });
  }

  findAll(): Promise<User[]> {
    return this.prismaService.user.findMany({});
  }

  findOne(id: number): Promise<User> {
    return this.prismaService.user.findUnique({ where: { id } });
  }

  update(id: number, data: Prisma.UserUpdateInput): Promise<User> {
    return this.prismaService.user.update({ where: { id }, data });
  }

  delete(id: number) {
    return this.prismaService.user.delete({ where: { id } });
  }
}
