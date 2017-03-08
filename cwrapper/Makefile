
CC=gcc

.DEFAULT=libmyc.so

libmyc.so : myc.c
	${CC} --shared -fPIC -o libmyc.so myc.c

clean:
	rm -rf libmyc.so

