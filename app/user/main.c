#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "user.c"
#include "userRepository.c"

int main()
{
   UserRepository repo;
   initRepository(&repo);

   User *alice = createUser(1, "Alice");
   User *bob = createUser(2, "Bob");

   addUser(&repo, alice);
   addUser(&repo, bob);

   User *searchedUser = findUserById(&repo, 1);
   if (searchedUser)
   {
      printf("Found user: %s\n", searchedUser->name);
   }
   else
   {
      printf("User not found.\n");
   }

   // Clean up memory
   free(alice);
   free(bob);

   return 0;
}