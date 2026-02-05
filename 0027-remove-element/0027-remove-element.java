class Solution {
    public int removeElement(int[] nums, int val) {

        int n = nums.length;
        for(int i=0; i<n; i++) {
            if (nums[i] == val) {
                if (nums[n - 1] != val){
                    nums[i] = nums[n - 1];
                    n--;
                } else {
                    n--;
                    i--;
                }
            } 
        }

        return n;
    }
}