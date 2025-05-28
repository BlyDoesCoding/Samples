#include <stdio.h>


/*prints the given string in cyan

*/
void info(char message[]) {

    printf(" \e[0;36m%s\e[0m\n", message);
    return;
}

/*prints the given string in RED with a ERROR before

*/
void error(char message[]) {

    printf("[\e[1;91mERROR\e[0m] \e[0;91m%s\e[0m\n", message);
    return;
}

/*prints the given string in yellow with a INFO before

*/
void warning(char message[]) {

    printf("[\e[1;93mINFO\e[0m] \e[0;93m%s\e[0m\n", message);
    return;
}
