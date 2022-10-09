from collections import deque


class Node:
    def __init__(self, name):
        self.sum_cost = float('inf')
        self.name = name
        self.points: list[tuple[Node, int]] = []
        self._visited: bool = False
        self.prev: Node = None

    def point_to(self, target_node: "Node", weight: int):
        self.points.append((target_node, weight))

    def __repr__(self):
        return self.name

    __str = __repr__

    def visit(self):
        self._visited = True

    def is_visited(self):
        return self._visited

    def __hash__(self):
        return self.name


n0 = Node("0")
n1 = Node("1")
n2 = Node("2")
n3 = Node("3")
n4 = Node("4")
n5 = Node("5")

n0.point_to(n1, 2)
n0.point_to(n2, 5)
n1.point_to(n3, 7)
n1.point_to(n2, 8)
n2.point_to(n3, 0)
n2.point_to(n4, 4)
n3.point_to(n5, 1)
n4.point_to(n3, 6)
n4.point_to(n5, 3)
n0.sum_cost = 0
queue = deque()
queue.append(n0)

while queue:
    node = queue.popleft()
    if node.is_visited():
        continue
    node.visit()
    for neighbor, cost in sorted(node.points, key=lambda item: item[1]):
        if neighbor.is_visited():
            continue
        if neighbor.sum_cost > cost + node.sum_cost:
            neighbor.sum_cost = cost + node.sum_cost
            neighbor.prev = node
        queue.append(neighbor)

print(f"To {n5.name}: {n5.sum_cost}")
prev_node = n5
while prev_node.prev:
    print(prev_node, end=f' <- ')
    prev_node = prev_node.prev
print(prev_node)
