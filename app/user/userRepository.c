#include <stdio.h>

// user.cの内容をインクルード
#include "user.c"

typedef struct
{
   User *users[100]; // 仮に最大100ユーザーとする
   int count;
} UserRepository;

UserRepository NewRepository()
{
   UserRepository repo;
   repo.count = 0;
   return repo;
}

int addUser(UserRepository *repo, User *user)
{
   if (repo->count >= 100)
   {
      return -1; // Repository is full
   }
   repo->users[repo->count++] = user;
   return 0;
}

User *findUserById(UserRepository *repo, int id)
{
   for (int i = 0; i < repo->count; i++)
   {
      if (repo->users[i]->id == id)
      {
         return repo->users[i];
      }
   }
   return NULL; // Not found
}
