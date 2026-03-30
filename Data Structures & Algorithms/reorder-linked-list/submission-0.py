# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        # [0 ~ 6]
        # length? 7
        # Brute Force
        # [2 4 6 8]
        # Time: O(n)
        # Space: O(n)
        nodes = []
        cur = head
        while cur is not None:
            nodes.append(cur)
            cur = cur.next
        
        # do this loop len() / 2 times
        cur = nodes[0]
        offset = 1
        offset2 = 1
        for i in range(len(nodes)-1):
            # even
            if i % 2 == 0:
                cur.next = nodes[len(nodes)-offset]
                offset += 1
                cur = cur.next
            # odd
            else:
                cur.next = nodes[offset2]
                offset2 += 1
                cur = cur.next

        cur.next = None
        return 
        