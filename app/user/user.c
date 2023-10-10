#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct
{
   int id;
   char name[50];
} User;

User *createUser(int id, const char *name)
{
   User *newUser = (User *)malloc(sizeof(User));
   if (!newUser)
   {
      exit(1);
   }
   newUser->id = id;
   strncpy(newUser->name, name, sizeof(newUser->name) - 1);
   newUser->name[sizeof(newUser->name) - 1] = '\0'; // Ensure null termination
   return newUser;
}