CC=gcc
CFLAGS=-I. -fPIC
OBJ=obj
SRC=Lyra2RE
TARGET=liblyra2rev2_hash.so

SOURCES := $(wildcard $(SRC)/*.c)
OBJS :=  $(patsubst $(SRC)/%.c, $(OBJ)/%.o, $(SOURCES)) $(OBJ)/Lyra2RE.o

$(TARGET): $(OBJS) 
	$(CC) -o $@ $^ $(CFLAGS) -shared 
	cp $@ /usr/lib

$(OBJ)/%.o: $(SRC)/%.c | $(OBJ)
	$(CC) -c -o $@ $< $(CFLAGS)

$(OBJ):
	mkdir obj

$(OBJ)/Lyra2RE.o: Lyra2RE.c
	$(CC) -c -o $@ $< $(CFLAGS)

.PHONY: clean
clean:
	rm -Rf $(OBJ) *.so 
