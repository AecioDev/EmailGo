package email

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	host     = "smtp.gmail.com"
	port     = 587
	user     = "espiranda.ms@gmail.com"
	password = "kqzwoyuzafphfoaw"
)

func EnviaEmail(fromEmail, toEmail, subjectEmail, bodyEmail string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", fromEmail)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subjectEmail)
	msg.SetBody("text/html", bodyEmail)

	dialer := gomail.NewDialer(host, port, user, password)

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	} else {
		return nil
	}
}

func GeneratePasswordRecoveryEmail(userName, userEmail, resetLink string) string {

	// Converte as imagens para base64
	sefazLogoBase64 := ImageToBase64("email/img/logo-sefaz-branca.png")
	cotinLogoBase64 := ImageToBase64("email/img/logo-cotin-branca.png")

	emailHTML := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Recuperação de Senha</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				margin: 0;
				padding: 0;
				color: #333;
			}
			.email-container {
				max-width: 600px;
				margin: 0 auto;
				background-color: #ffffff;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
			.header {
				background-color: #004F9F;
				padding: 20px;
				display: flex;
				align-items: center;
			}
			.header img {
				width: 150px; 
				height: auto;
			}
			.content {
				margin-top: 20px;
				font-size: 16px;
				line-height: 1.5;
			}
			.content a {
				color: #0066cc;
				text-decoration: none;
			}
			.prefooter {
				margin-top: 30px;
				font-size: 12px;
				color: #777;
				text-align: center;
			}
			.footer {
				background-color: #004F9F;
				padding: 20px;
				display: flex;
                justify-content: center;
			}                     
			.footer img {
				width: 150px; 
				height: auto; 
			}
			.posfooter {
				background-color: #0094EA;
				width: auto;
				height: 18px;
				padding: 10px;
				font-size: 12px;
				color: #fff;
				text-align: center;
                display: flex;
                justify-content: center;
				align-items: center;
			}
		</style>
	</head>
	<body>
		<div class="email-container">
			<!-- Cabeçalho com fundo azul e logo -->
			<div class="header">
				<img src="%s" alt="logo SEFAZ" style="width: 150px; height: auto;">
			</div>

			<!-- Saudação e conteúdo principal -->
			<div class="content">
				<p>Olá, <strong>%s</strong>,</p>
				<p>Você solicitou a recuperação de sua senha. Para redefinir sua senha, clique no link abaixo:</p>
				<p><a href="%s">Redefinir senha</a></p>
				<p>Se o link não funcionar, copie e cole o seguinte endereço no seu navegador:</p>
				<p>%s</p>
			</div>

			<!-- Rodapé -->
			<div class="prefooter">
				<p>Este e-mail foi enviado para <strong>%s</strong>.</p>
				<p>Se você não solicitou a recuperação de senha, por favor, desconsidere este e-mail.</p>
			</div>
            <!-- Rodapé com fundo azul e logo -->
			<div class="footer">
				<img src="%s" alt="logo COTIN" style="width: 150px; height: auto;">
			</div>
            <div class="posfooter">
        <p>2024 COTIN - Todos os direitos reservados</p>
      </div>
		</div>
	</body>
	</html>
	`, sefazLogoBase64, userName, resetLink, resetLink, userEmail, cotinLogoBase64)

	return emailHTML
}

// Converte a imagem para base64
func ImageToBase64(imagePath string) string {
	// Lê o arquivo da imagem
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de imagem: %v", err)
	}

	// Converte os bytes da imagem para base64
	base64Encoding := base64.StdEncoding.EncodeToString(imageBytes)

	// Retorna a string no formato correto para o HTML
	return fmt.Sprintf("data:image/png;base64,%s", base64Encoding)
}
