package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

func sshConnect(user string, host string, port int, password string) error {
	// Configurer l'authentification par nom d'utilisateur et mot de passe
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Se connecter au serveur SSH
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
	if err != nil {
		return fmt.Errorf("Erreur lors de la connexion au serveur SSH: %s", err)
	}
	defer conn.Close()

	// Ouvrir une session SSH
	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ouverture de la session SSH: %s", err)
	}
	defer session.Close()

	// Configurer les entrées/sorties standard
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	// Désactiver l'affichage des messages de bienvenue et des invites de commande
	session.RequestPty("xterm", 80, 40, ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	})

	// Démarrer le shell interactif
	err = session.Shell()
	if err != nil {
		return fmt.Errorf("Erreur lors du démarrage du shell: %s", err)
	}

	// Effacer la console
	fmt.Print("\033[H\033[2J")

	// Attendre la fin de la session
	err = session.Wait()
	if err != nil {
		return fmt.Errorf("Erreur lors de l'exécution de la commande: %s", err)
	}

	return nil
}

func main() {
	// Charger les variables d'environnement à partir du fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %s", err)
	}

	// Vérifier si les variables d'environnement existent
	user, exists := os.LookupEnv("USER")
	if !exists {
		log.Fatal("La variable d'environnement USER n'existe pas")
	}

	host, exists := os.LookupEnv("HOST")
	if !exists {
		log.Fatal("La variable d'environnement HOST n'existe pas")
	}

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		log.Fatal("La variable d'environnement PASSWORD n'existe pas")
	}

	portStr, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("La variable d'environnement PORT n'existe pas")
	}

	// Convertir la variable d'environnement PORT en entier
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion de la variable d'environnement PORT en entier : %s", err)
	}

	// Se connecter au serveur SSH
	err = sshConnect(user, host, port, password)
	if err != nil {
		log.Fatalf("Erreur lors de la connexion SSH: %s", err)
	}
}
