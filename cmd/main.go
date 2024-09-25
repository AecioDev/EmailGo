package main

import (
	"enviaemail/email"
	"fmt"
)

func main() {
	fmt.Println("Enviando E-mail")

	// Exemplo de uso
	userName := "Aécio"
	userEmail := "espiranda@hotmail.com"
	resetLink := "https://example.com/redefinir-senha"

	emailHTML := email.GeneratePasswordRecoveryEmail(userName, userEmail, resetLink)
	fmt.Println(emailHTML)

	err := email.EnviaEmail("aemiranda@fazenda.ms.gov.br", userEmail, "Teste Email GO", emailHTML)
	if err != nil {
		fmt.Println(err)
	}
}
