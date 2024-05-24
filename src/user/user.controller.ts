import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { UserService } from './user.service';
import { CreateUserDto, UpdateUserDto } from './dto/user';
import { User } from '@prisma/client';

@Controller('user')
export class UserController {
  constructor(
    private readonly userService: UserService,
  ) { }

  @Post()
  async create(@Body() createUserDto: CreateUserDto): Promise<User> {
    return this.userService.create(createUserDto);
  }

  @Get()
  async findAll(): Promise<User[]> {
    return this.userService.findAll();
  }

  @Get(':id')
  async findOneById(@Param('id') id: string): Promise<User> {
    return this.userService.findOneById(+id);
  }

  @Get(':email')
  async findOneByEmail(@Param('email') email: string): Promise<User> {
    return this.userService.findOneByEmail(email);
  }


  @Patch(':id')
  async update(@Param('id') id: string, @Body() updateUserDto: UpdateUserDto) {
    return this.userService.update(+id, updateUserDto);
  }

  @Delete(':id')
  async remove(@Param('id') id: string) {
    return this.userService.delete(+id);
  }

  @Post()
  async resendOtp(@Body() data: { email: string }): Promise<string | Error> {
    return this.userService.resendOtp(data.email);
  }

  @Post()
  async verifyOtp(@Body() data: { email: string, otp: string }): Promise<string | Error> {
    return this.userService.verifyOtp(data.email, data.otp);
  }
}
