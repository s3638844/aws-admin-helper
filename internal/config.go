package internal

type AppConfig struct {
	Region     string `yaml:"region"`
	Credential AWSCredential `yaml:"credential"`

}

type AWSCredential struct {
	AccessKey		string	`yaml:"AWS_ACCESS_KEY_ID"`
	SecretKey		string	`yaml:"AWS_SECRET_ACCESS_KEY"`
	SessionToken	string	`yaml:"AWS_SESSION_TOKEN"`
}

type TerraformVars struct {
	Region		string	`yaml:"AWS_DEFAULT_REGION"`
	AdPassword 	string	`yaml:"TF_VAR_aws_active_directory_password"`
}

func promptStartUrl