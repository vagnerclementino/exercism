#ifndef EXERCISE_NAME_H
#define EXERCISE_NAME_H

typedef struct tm {
   int tm_year;
   int tm_mon;
   int tm_mday;
   int tm_hour;
   int tm_min;
   int tm_sec;
   int tm_isdst;
} time_t;

time_t mktime(time_t*);

time_t gigasecond_after(time_t );

#endif /* ifndef EXERCISE_NAME_H*/
