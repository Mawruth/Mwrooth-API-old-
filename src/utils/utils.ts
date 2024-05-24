import * as fs from "fs";

export function generateOtpVerificationMessage(otp: string): string {
	const msg = fs.readFileSync(__dirname + '../templates/otp_verify.txt', 'utf8');
	return msg.replace('{otp}', otp);
}