class Twofer {

    private static final String ANSWER_TEMPLATE = "One for {NAME}, one for me.";

    private static final String NAME_PATTERN = "{NAME}";

    private static final String NO_NAME_VALUE = "you";

    String twofer(String name) {

        String twoferAnswer = null;

        if (name == null || name.trim().isEmpty()) {
            twoferAnswer = ANSWER_TEMPLATE.replace(NAME_PATTERN, NO_NAME_VALUE);
        }
        else{

            twoferAnswer = ANSWER_TEMPLATE.replace(NAME_PATTERN, name);
        }

        return twoferAnswer;
    }
}
