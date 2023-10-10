#include <stdio.h>
#include <stdlib.h>

// ノードの定義
typedef struct Node
{
   int data;          // データ部分
   struct Node *next; // 次のノードへのポインタ部分
} Node;

// ノードの生成
Node *createNode(int data)
{
   Node *newNode = (Node *)malloc(sizeof(Node));
   if (newNode == NULL)
   {
      exit(1); // メモリ確保に失敗
   }
   newNode->data = data;
   newNode->next = NULL;
   return newNode;
}

// ノードの末尾に追加
void append(Node **head, int data)
{
   Node *newNode = createNode(data);
   if (*head == NULL)
   {
      *head = newNode;
   }
   else
   {
      Node *temp = *head;
      while (temp->next != NULL)
      {
         temp = temp->next;
      }
      temp->next = newNode;
   }
}

// リンクリストの全ノードを表示
void printList(Node *head)
{
   Node *temp = head;
   while (temp != NULL)
   {
      printf("%d -> ", temp->data);
      temp = temp->next;
   }
   printf("NULL\n");
}

// リンクリストのメモリを解放
void freeList(Node *head)
{
   Node *temp;
   while (head != NULL)
   {
      temp = head;
      head = head->next;
      free(temp);
   }
}

int main()
{
   Node *head = NULL;

   append(&head, 1);
   append(&head, 2);
   append(&head, 3);
   append(&head, 4);
   append(&head, 5);

   printList(head);
   freeList(head);

   return 0;
}
