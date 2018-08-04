import java.util.stream.IntStream;

class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {

        return (int) Math.pow((input *(input + 1)/2),2);

    }

    int computeSumOfSquaresTo(int input) {
        int sumOfSquares = 1;

        sumOfSquares = IntStream
                       .rangeClosed(1,input)
                       .map(i ->i*i)
                       .sum();

        return sumOfSquares;

    }


    int computeDifferenceOfSquares(int input) {
        return ( computeSquareOfSumTo(input) -
                 computeSumOfSquaresTo(input)
                );
    }

}
