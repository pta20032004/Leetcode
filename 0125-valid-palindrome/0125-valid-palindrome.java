class Solution {
    public boolean isPalindrome(String s) {
        int p1 = 0;
        int p2 = s.length() - 1;

        for (int i = 0; i < s.length(); i++) {
            if (p1 >= p2) return true;
            if (!Character.isLetter(s.charAt(p2)) && !Character.isDigit(s.charAt(p2))) {
                p2--;
                continue;
            } 

            if (!Character.isLetter(s.charAt(p1)) && !Character.isDigit(s.charAt(p1))) {
                p1++;
                continue;
            }

            if (Character.toLowerCase(s.charAt(p1)) != Character.toLowerCase(s.charAt(p2))) return false; 
            p1++;
            p2--;
        }
        return true;
    }
}