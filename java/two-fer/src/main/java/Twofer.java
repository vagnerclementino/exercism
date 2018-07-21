class Twofer {

    private static final String ANSWER_TEMPLATE = "One for %s, one for me.";
    private static final String NO_NAME_VALUE = "you";

    String twofer(String name) {

        String twoferAnswer = null;

        if (name == null || name.trim().isEmpty()) {
            twoferAnswer = String.format(ANSWER_TEMPLATE, NO_NAME_VALUE);
        }
        else{

            twoferAnswer = String.format(ANSWER_TEMPLATE, name);
        }

        return twoferAnswer;
    }
}
