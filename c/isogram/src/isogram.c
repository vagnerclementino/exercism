#include "isogram.h"
#include <stdio.h>

bool is_valid_character(const char c){
   bool is_valid = false;
   switch (c) {
       case '-':
           is_valid = true;
           break;
       case ' ':
           is_valid = true;
           break;
       default:
           is_valid = false;
   }
   return is_valid;
}

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
                    if (!is_valid_character(phrase[i])) {
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





