echo "This script will populate database with data extracted from wahapedia CSV: https://wahapedia.ru/wh40k9ed/the-rules/data-export/"
echo "!! It will drop database and repopulate it !!"
read -p "Are you sure you want to continue? y/n " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    chmod +x lib/populate/populate.go
    go run lib/populate/populate.go
fi
