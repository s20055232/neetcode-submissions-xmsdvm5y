# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        # At the beginning, cur = None
        head = None
        
        if list1 is None:
            return list2
        
        if list2 is None:
            return list1
        
        if list1.val > list2.val:
            head = list2
            list2 = list2.next
        else:
            head = list1
            list1 = list1.next
        
        cur = head
        
        # for loop this until list1 and list2 become None
        while list1 is not None or list2 is not None:
            if list1 is None:
                cur.next = list2
                list2 = list2.next
            elif list2 is None:
                cur.next = list1
                list1 = list1.next
            elif list1.val > list2.val:
                cur.next = list2
                list2 = list2.next
            else:
                cur.next = list1
                list1 = list1.next
            
            cur = cur.next
        cur.next = None
        return head