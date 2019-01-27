#include<stdio.h>
#include<stdlib.h>
/*
	图
*/
struct BNode
{
	int data;
	struct BNode *next;
};

struct FNode
{
	char ch;
	struct BNode *first_node;
};

struct Graph
{
	int vernum;
	int sidnum;
	struct FNode* node[100];
};

int getPosition(char ch, struct Graph *g){
	int i = 0;
	for(i = 0; i < g->vernum; i++){
		if(g->node[i]->ch == ch){
			return i;
		}
	}
	return -1;
}

void insertNode(struct BNode *g, struct BNode *node){
	struct BNode *temp = g;
	while(temp->next){
		temp = temp->next;
	}
	temp->next = node;
	printf("被插入数据：%d\n", temp->next->data);
}

struct Graph* createGraph(){
	int i = 0;
	struct Graph *g = (struct Graph*)malloc(sizeof(struct Graph));
	if(g == NULL){
		printf("内存空间不足。\n");
		return NULL;
	}
	printf("请输入顶点数：\n");
	scanf("%d", &g->vernum);
	printf("请输入边数：\n");
	scanf("%d", &g->sidnum);
	printf("请输入顶点元素：\n");
	for(i = 0; i < g->vernum; i++){
		getchar();
		g->node[i] = (struct FNode*)malloc(sizeof(struct FNode));
		if(g->node[i] == NULL){
			printf("内存空间不足。\n");
			return NULL;
		}
		g->node[i]->first_node = NULL;
		scanf("%c", &g->node[i]->ch);
		// printf("here!!!: %c\n", g->node[i]->ch);
	}
	// for(i = 0; i < g->vernum; i++){
	// 	printf("%c\n", g->node[i]->ch);
	// }
	printf("请输入顶点之间的关系：（即相连顶点，以逗号分隔）\n");
	for(i = 0; i < g->sidnum; i++){
		char ch1, ch2;
		getchar();
		scanf("%c,%c", &ch1, &ch2);
		printf("%c,%c\n", ch1,ch2);
		struct BNode *node1, *node2;
		node1 = (struct BNode*)malloc(sizeof(struct BNode));
		if(node1 == NULL){
			printf("内存空间不足。\n");
			return NULL;
		}
		node2 = (struct BNode*)malloc(sizeof(struct BNode));
		if(node2 == NULL){
			printf("内存空间不足。\n");
			return NULL;
		}
		node1->data = getPosition(ch1, g);
		node2->data = getPosition(ch2, g);
		node1->next = NULL;
		node2->next = NULL;
		if(g->node[node1->data]->first_node == NULL){
			g->node[node1->data]->first_node = (struct BNode*)malloc(sizeof(struct BNode));
			if(g->node[node1->data]->first_node == NULL){
				printf("内存空间不足。\n");
				return NULL;
			}
			g->node[node1->data] -> first_node = node2;
			printf("被插入数据：%d\n", g->node[node1->data]->first_node->data);
		}else{
			insertNode(g->node[node1->data]->first_node, node2);
		}
	}
	return g;
}

void printGraph(struct Graph *g){
	int i;
	struct BNode *n;
	for(i = 0; i < g->vernum; i++){
		if(g->node[i]->first_node == NULL)
			printf("%c\n", g->node[i]->ch);
		else
			printf("%c -> ", g->node[i]->ch);
		n = g->node[i]->first_node;
		while(n != NULL){
			if(n -> next != NULL)
				printf("%d -> ", n->data);
			else
				printf("%d\n", n->data);
			n = n->next;
		}
	}
}

int main(){
	struct Graph *g;
	g = createGraph();
	printGraph(g);
	return 0;
}




