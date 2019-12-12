#include <stdio.h> 
#include <string.h> 
#include <fcntl.h> 
#include <sys/stat.h> 
#include <sys/types.h> 
#include <unistd.h> 

const char *pipeFile = "go-c-pipe.pipe";

int main(int argc, char **argv) {

    int fd = open(pipeFile, O_WRONLY);

    int buf_size = 1024;
    char buf[buf_size];

    while(1) {
        fgets(buf, buf_size, stdin);
        write(fd, buf, strlen(buf)+1);
    }

    return 0;
}