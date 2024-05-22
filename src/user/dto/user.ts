export class CreateUserDto {
	fullName: string;
	userName: string;
	email: string;
	password: string;
	avatar?: string;
}

export class UpdateUserDto {
	fullName?: string;
	userName?: string;
	email?: string;
	password?: string;
	avatar?: string;
}