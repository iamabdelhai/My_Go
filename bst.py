class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class BST:
    def __init__(self):
        self.root = None

    def insert(self, val):
        if self.root is None:
            self.root = Node(val)
            return
        current = self.root
        while True:
            if val < current.val:
                if current.left is None:
                    current.left = Node(val)
                    return
                current = current.left
            else:
                if current.right is None:
                    current.right = Node(val)
                    return
                current = current.right

    def search(self, val):
        current = self.root
        while current:
            if val == current.val:
                return True
            elif val < current.val:
                current = current.left
            else:
                current = current.right
        return False

    def delete(self, val):
        self.root = self._delete(self.root, val)

    def _delete(self, node, val):
        if not node:
            return None
        if val < node.val:
            node.left = self._delete(node.left, val)
        elif val > node.val:
            node.right = self._delete(node.right, val)
        else:
            if not node.left:
                return node.right
            elif not node.right:
                return node.left
            successor = self._find_min(node.right)
            node.val = successor.val
            node.right = self._delete(node.right, successor.val)
        return node

    def _find_min(self, node):
        while node.left:
            node = node.left
        return node

    def in_order(self):
        result = []
        self._in_order(self.root, result)
        return result

    def _in_order(self, node, result):
        if node:
            self._in_order(node.left, result)
            result.append(node.val)
            self._in_order(node.right, result)

    def pre_order(self):
        result = []
        self._pre_order(self.root, result)
        return result

    def _pre_order(self, node, result):
        if node:
            result.append(node.val)
            self._pre_order(node.left, result)
            self._pre_order(node.right, result)

    def height(self):
        return self._height(self.root)

    def _height(self, node):
        if not node:
            return 0
        left_h = self._height(node.left)
        right_h = self._height(node.right)
        return max(left_h, right_h) + 1

if __name__ == "__main__":
    bst = BST()
    values = [50, 30, 70, 20, 40, 60, 80]
    print("Inserting values:", values)
    for v in values:
        bst.insert(v)
    print("\n── Traversals ──")
    print("In-Order  (sorted):", bst.in_order())
    print("Pre-Order         :", bst.pre_order())
    print("Tree Height       :", bst.height())
    print("\n── Search ──")
    print("Search(40) →", bst.search(40))
    print("Search(99) →", bst.search(99))
    print("\n── Delete 30 ──")
    bst.delete(30)
    print("In-Order after delete:", bst.in_order())