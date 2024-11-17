package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {
	// Caminho para o arquivo JSON da chave privada do Firebase
	opt := option.WithCredentialsFile("C:/Users/caioo/Downloads/totp-mobile-fire-firebase-adminsdk-8sbfk-9879ea7ca3.json")

	// Inicializa o app Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase App: %v", err)
	}

	// Obtém o cliente de mensagens
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Erro ao inicializar o cliente do Firebase Messaging: %v", err)
	}

	// Enviar notificação push
	sendPushNotification(client)
}

func sendPushNotification(client *messaging.Client) {
	// Definir a mensagem
	message := &messaging.Message{
		Topic: "login_validation", // O tópico para o qual enviar
		Notification: &messaging.Notification{
			Title: "Validação de Login",
			Body:  "Por favor, confirme sua tentativa de login no aplicativo.",
		},
		Data: map[string]string{
			"userID":    "123456",
			"sessionID": "abcde12345",
			"ipAddress": "192.168.1.1",
		},
	}

	// Log para visualizar a mensagem sendo enviada
	log.Printf("Enviando mensagem: %+v", message)

	// Enviar a mensagem
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}
	log.Printf("Mensagem enviada com sucesso: %s", response)
}
