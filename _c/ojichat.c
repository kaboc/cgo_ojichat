#include <stdio.h>
#include "libojichat.h"

int main() {
    setvbuf(stdout, NULL, _IONBF, 0);

    char str[256];
    printf("お名前： ");
    scanf("%s", str);

    printf(getMessage(str));

    return 0;
}
