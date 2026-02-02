/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
// import java.util.Stack;
// class Solution {
//     public ListNode reverseList(ListNode head) {
//         if (head == null) {
//             return head;
//         }
//         Stack<ListNode> stack = new Stack<>();
//         ListNode cur = head;
//         while (cur.next != null) {
//             stack.push(cur);
//             ListNode prev = cur;
//             cur = cur.next;
//             prev.next = null;
//         }

//         head = cur;
//         while (!stack.isEmpty()) {
//             cur.next = stack.pop();
//             cur = cur.next;
//         }
//         return head;
//     }
// }

// Toi uu
class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null) {
            return head;
        }

        if (head.next == null) {
            // head.next.next = head;
            // ListNode temp = head.next.next;
            // head.next = null;
            // return temp;
            return head;
        }

        ListNode cur = head.next;
        ListNode prev = head;
        ListNode tempHead = head;
        ListNode future;
        while (cur != null) {
            future = cur.next;
            cur.next = prev;
            prev = cur;
            cur = future;
        }

        tempHead.next = null;
        return prev;

    }
}