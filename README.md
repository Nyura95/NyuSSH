# NyuSSH
NyuSSH est un programme en Go qui permet de se connecter automatiquement à un serveur SSH en utilisant un nom d'utilisateur et un mot de passe.

## Installation
Assurez-vous que Go est installé sur votre système. Vous pouvez télécharger Go à partir du site officiel: https://golang.org/dl/

Clonez le dépôt Git sur votre système:

```
git clone https://github.com/Nyura95/NyuSSH.git
```

Installez les dépendances en exécutant la commande suivante:

```
go get github.com/joho/godotenv
go get golang.org/x/crypto/ssh
```

Créez un fichier ".env" à la racine de votre projet et ajoutez-y les variables d'environnement suivantes:

```
USERNAME=votre_nom_d_utilisateur
PASSWORD=votre_mot_de_passe
HOST=adresse_ip_du_serveur_ssh
PORT=port_du_serveur_ssh
```

Exécutez le programme en exécutant la commande suivante:
```
go run main.go
```

## Utilisation
Le programme se connectera automatiquement au serveur SSH en utilisant les informations d'identification fournies dans le fichier ".env". Une fois connecté, vous pourrez saisir des commandes dans le terminal.

Licence
Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus d'informations.

Auteur
Ce projet a été créé par Nyura.