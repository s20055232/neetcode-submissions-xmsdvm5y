# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head is None:
            return False
        
        fast, slow = head, head
        while True:
            slow = slow.next
            if slow is None:
                return False

            fast = fast.next
            if fast is None:
                return False

            fast = fast.next
            if fast is None:
                return False
            
            if slow.val == fast.val:
                return True

        return False
