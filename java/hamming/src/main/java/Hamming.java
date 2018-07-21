import java.util.stream.IntStream;


class Hamming {

    private String leftStrand;
    private String rightStrand;
    private int lengthStrand;

    Hamming(String leftStrand, String rightStrand) {

        /*
         * Validade if the length two strings are sames.
         */
         if (leftStrand.length() != rightStrand.length()) {
             throw new IllegalArgumentException("leftStrand and rightStrand " +
                                                "must be of equal length.");
         }
        else{
            this.leftStrand = leftStrand;
            this.rightStrand = rightStrand;
            this.lengthStrand = leftStrand.length();
        }
    }

    public int getHammingDistance() {

         int hammingDistance = 0;
         /*
          * It interate throught a range between 0 and the length of the
          *'lengthStrand'. Next, we filter the cases where the characters are
          * different. Finally we count that cases.
          */
         hammingDistance = (int) IntStream
                                    .range(0, this.lengthStrand)
                                    .filter(i -> this.leftStrand.charAt(i) !=
                                                 this.rightStrand.charAt(i)
                                           )
                                    .count();
         return hammingDistance;
    }
}
