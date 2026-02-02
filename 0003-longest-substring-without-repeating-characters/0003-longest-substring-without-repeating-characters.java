import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s.length() == 0) {
            return 0;
        }

        String start = s.substring(0, 1);
        int curLength = 0;
        int maxLength = 0;

        Queue<String> queue = new ArrayDeque<>();
        for (int i = 0; i < s.length(); i++) {
            String cur = s.substring(i, i + 1);

            if (queue.contains(cur)) {
                helper(queue, cur);
                queue.add(cur); 
            } else {
                queue.add(cur);
            }

            if (queue.size() > maxLength) {
                maxLength = queue.size();
            }
        }

        return maxLength;
    }

    public void helper(Queue<String> queue, String a) {
        while (true) {
            String temp = queue.poll();
            if (temp.equals(a)) {
                break;
            }
        }
    }
}