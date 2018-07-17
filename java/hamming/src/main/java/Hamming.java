class Hamming {

    private String leftStrand;
    private String rightStrand;
    private int lengthStrand;

    Hamming(String leftStrand, String rightStrand) {

        /*
         * Validade if the length two strings are sames.
         */
         if (leftStrand.length() != rightStrand.length()) {
             throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
         }
        else{
            this.leftStrand = leftStrand;
            this.rightStrand = rightStrand;
            this.lengthStrand = leftStrand.length();
        }
    }

    public int getHammingDistance() {

         int hammingDistance = 0;

         for (int i = 0; i < this.lengthStrand; i++) {
             if (this.leftStrand.charAt(i) != this.rightStrand.charAt(i)) {
                 hammingDistance++;
             }
         }
         return hammingDistance;
    }
}
