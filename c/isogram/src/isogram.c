#include "isogram.h"
#include <stdio.h>

bool is_isogram(const char phrase[]){

    /* Variable that define if parameter
     * phrase is
     * isogram */

    bool is_isogram = true;
    int i = 0;
    int j = 0;
    /* Run through variable phrase and
     * exit when the character in position *i*
     * is equal the character in position *j*
     */
    while (phrase[i] != '\0'){
        j = i + 1;
        while(phrase[j] != '\0'){
            if (i != j ) {
                if (tolower(phrase[i]) == tolower(phrase[j]))
                {
                    if (phrase[i] != '-' && phrase[i] != ' ') {
                        is_isogram = false;
                        break;
                    }
                }
            }
            j++;
        }
        /* If a repeating letter is found 
         * exit from the loop, otherwise, add one
         * to variable i
         * */
        if (is_isogram) {
            i++;
        }else
        {
         break;
        }

    }
    return is_isogram;

}





