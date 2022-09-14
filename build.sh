#/bin/bash
rm Hello.class libHello.h libHello.so
#linux可能需要插件mod.
#go mod init a
#go mod tidy
go build -buildmode=c-shared -o libHello.so
javac Hello.java
java -Djava.library.path=. Hello
