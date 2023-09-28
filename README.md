# GoExo1


Description
Ce projet Go permet de récupérer et de trier les répertoires GitHub d'un utilisateur par date de mise à jour. Il effectue également des opérations git fetch et git pull sur chaque répertoire cloné en local, et enfin, il archive tous ces répertoires dans un fichier zip.

## Fonctionnalités :

Récupération des répertoires d'un utilisateur depuis l'API GitHub.  
Tri des répertoires par date de mise à jour.  
Clonage des répertoires en local.  
Exécution des commandes git fetch et git pull sur les répertoires clonés.  
Archivage des répertoires clonés en un fichier zip.  

### Configuration

Avant de lancer le programme, assurez-vous de configurer un token d'accès personnel GitHub comme variable d'environnement.


### Utilisation

Clonez ce répertoire dans votre espace de travail Go.  
Naviguez jusqu'au répertoire du projet.  
Exécutez la commande go run main.go.  

## Structure du Projet :

getRepo.go: Contient la logique pour récupérer les informations des répertoires depuis l'API GitHub.   
csvfile.go: Contient la logique pour écrire les informations des répertoires dans un fichier CSV.  
archive.go: Contient la logique pour archiver les répertoires clonés en un fichier zip.  
git.go: Contient la logique pour effectuer des opérations git fetch et git pull sur les répertoires clonés.  
main.go: Point d'entrée du programme, contient la logique principale de coordination.  
