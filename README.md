/************************************************/
/*                         GoFractals - Go et Julia                          */
/************************************************/


1. Comment compiler et lancer notre projet Go :

  a. Le reposiroty svn est : http://xp-dev.com/svn/Go_svn/
  b. Un dossier Go_svn contient toute l'arborescence
  c. Faites un svn up si vous aviez déjà le dossier pour avoir la dernière version du svn
  d. Votre .bashrc doit contenir les variables suivantes :
    export GOWORKSPACE=$HOME/Go_svn
    export GOPATH=$GOWORKSPACE/myfractale:$GOWORKSPACE/myhttpserver:$GOWORKSPACE/termbox-go
    export GOROOT=/usr/local/go/
    export PATH=$PATH:$GOROOT/bin
  e. Allez dans le dossier fracture : $ cd Go_svn/myfractale/src/fracture/
  f. Lancer la commande : $ go install fracture
  g. Allez dans le dossier myhttpdserver : $ cd Go_svn/myhttpserver/src/myhttpserver/
  h. Lancez la commande : $ go build myhttpdserver
  i. Lancez le serveur : $ ./myhttpdserver
  j. Utilisez votre navigateur pour générer les fractales : localhost:8080


2. Utiliser l'interface Web :

  a. Le paramètre C peut-être définit avec les variables a et b
  b. Lancer la génération en appuyant sur le bouton 'Go'
  c. Zoomez et déplacez-vous dans l'image avec les boutons d'actions

