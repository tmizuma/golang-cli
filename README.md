rm -rf /Users/tamizuma/Desktop/f1/
rm -rf /Users/tamizuma/Desktop/f2/

mkdir /Users/tamizuma/Desktop/f1/
mkdir /Users/tamizuma/Desktop/f2/
mkdir /Users/tamizuma/Desktop/f2/col10

touch /Users/tamizuma/Desktop/f1/0001.png
touch /Users/tamizuma/Desktop/f1/0002.png
touch /Users/tamizuma/Desktop/f1/0003.png
touch /Users/tamizuma/Desktop/f1/0004.png
touch /Users/tamizuma/Desktop/f1/0005.png
touch /Users/tamizuma/Desktop/f1/0006.png
touch /Users/tamizuma/Desktop/f1/0007.png
touch /Users/tamizuma/Desktop/f1/0008.png
touch /Users/tamizuma/Desktop/f1/0009.png

go run main.go format -inputFolder /Users/tamizuma/Desktop/f1 -outputFolder /Users/tamizuma/Desktop/f2 -chunkSize 3
