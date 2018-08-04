import java.util.HashMap;


class Scrabble {

    private HashMap<String, Integer> mapScore;
    private String word;

    Scrabble(String word) {

        this.word = word.trim().toUpperCase();

        this.mapScore = new HashMap<String, Integer>();

        this.mapScore.put("A", 1);
        this.mapScore.put("E", 1);
        this.mapScore.put("I", 1);
        this.mapScore.put("O", 1);
        this.mapScore.put("U", 1);
        this.mapScore.put("L", 1);
        this.mapScore.put("N", 1);
        this.mapScore.put("R", 1);
        this.mapScore.put("S", 1);
        this.mapScore.put("T", 1);

        this.mapScore.put("D", 2);
        this.mapScore.put("G", 2);

        this.mapScore.put("B", 3);
        this.mapScore.put("C", 3);
        this.mapScore.put("M", 3);
        this.mapScore.put("P", 3);

        this.mapScore.put("F", 4);
        this.mapScore.put("H", 4);
        this.mapScore.put("V", 4);
        this.mapScore.put("W", 4);
        this.mapScore.put("Y", 4);

        this.mapScore.put("K",5);

        this.mapScore.put("J", 8);
        this.mapScore.put("X", 8);

        this.mapScore.put("Q", 10);
        this.mapScore.put("Z", 10);

    }

    int getScore() {

        int score = 0;
        try {
            for (String ch: this.word.split("")) {
                 if (this.mapScore.containsKey(ch)){
                    score += this.mapScore.get(ch);
                 }
            }
        } catch(Exception e){
            e.printStackTrace();
        }
        return score;
    }
}
