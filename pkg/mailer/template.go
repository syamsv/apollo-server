package mailer

import "fmt"

var AccountActivation = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Account Activation</title>
    <link rel="stylesheet" href="activation.css">
	<style>
	body {
		font-family: Arial, sans-serif;
		background-color: #f0f0f0;
		margin: 0;
		padding: 0;
	}
	.h-16{
		width:60px;
	  }
	.container {
		max-width: 600px;
		margin: 30px auto;
		padding: 20px;
		background-color: #fff;
		border: 1px solid #ccc;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		text-align: center;
	}
	
	h1 {
		color: #333;
	}
	
	.activation-btn {
		margin-top: 20px;
	}
	
	.activation-btn a {
		display: inline-block;
		padding: 10px 20px;
		background-color: #007bff;
		color: #fff;
		text-decoration: none;
		border-radius: 5px;
	}
	
	.activation-btn a:hover {
		background-color: #0056b3;
	}
	
	p {
		color: #666;
		line-height: 1.6;
	}
	
	</style>
</head>
<body>
    <div class="container">
	<svg class="h-16" viewBox="0 0 700 700" xmlns="http://www.w3.org/2000/svg">
    <path fill="#2563eb"
        d="M352.387 27.657c-142.618 0-214.9 73.675-230.664 152.376-9.03 45.08.511 91.966 28.217 128.604a78.428 78.428 0 0 1 43.722-30.408c41.923-11.228 84.357 12.439 96.936 48.524l2.03 5.824-3.297 5.215c-63.518 100.401-89.642 188.902-37.485 306.334l43.344-71.365c-17.521-66.234-17.472-138.467 10.99-204.372l23.457 10.129c-38.612 89.39-19.25 194.999 25.935 274.911 47.369-83.363 57.785-194.6 29.274-274.764l24.073-8.561c21.539 60.578 23.219 134.246 5.754 203.602l42.77 70.42c52.15-117.432 26.04-205.933-37.485-306.334l-3.297-5.215 2.03-5.824c10.227-29.323 40.152-50.449 73.5-51.31a82.236 82.236 0 0 1 23.436 2.786 78.239 78.239 0 0 1 37.1 22.323c21.147-34.685 29.148-77.448 22.239-118.867C561.68 102.081 494.914 27.657 352.38 27.657zM213.136 301.14a52.283 52.283 0 0 0-12.873 1.778 53.116 53.116 0 0 0-37.695 65.289c7.343 27.384 34.328 43.68 61.712 38.283 10.486-25.592 24.038-50.897 39.76-76.3-8.456-16.233-28.945-29.358-50.911-29.043zm283.01 0c-21.959-.315-42.455 12.81-50.911 29.043 15.722 25.403 29.274 50.708 39.76 76.3 27.384 5.404 54.376-10.892 61.712-38.276a53.123 53.123 0 0 0-37.695-65.289 52.318 52.318 0 0 0-12.873-1.778zM176.904 423.787c-19.019 53.571-27.272 110.691-25.942 163.534 17.276-13.993 33.663-29.925 48.923-47.201-1.911-36.981 3.556-72.338 14.756-106.96a79.107 79.107 0 0 1-37.737-9.373zm354.529.497a79.128 79.128 0 0 1-36.792 8.876c11.326 35.007 16.8 70.756 14.7 108.178 14.959 16.807 30.982 32.312 47.852 45.976 1.33-52.675-6.874-109.613-25.767-163.037z" />
</svg>
        <h1>Account Activation</h1>
        <p>Dear User,</p>
        <p>Thank you for signing up with us. To activate your account, please click the button below:</p>
        <div class="activation-btn">
            <a href="http://localhost:3000/uas/activate/%s">Activate Account</a>
        </div>
        <p>If you did not sign up for an account on our website, please ignore this email.</p>
        <p>Best regards,</p>
        <p>Apollo Service</p>
    </div>
</body>
</html>
`

func ReturnHtmlTemplate(id string) string {
	return fmt.Sprintf(AccountActivation, id)
}
