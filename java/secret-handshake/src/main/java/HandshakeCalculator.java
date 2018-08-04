import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class HandshakeCalculator {

	private static final int REVERSE = 16;
	private static final int JUMP = 8;
    private static final int CLOSE_YOUR_EYES = 4;
    private static final int DOUBLE_BLINK = 2;
    private static final int WINK = 1;


    List<Signal> calculateHandshake(int number) {

        boolean reverseList = false;
        List<Signal> handShakeList = new ArrayList<Signal>();

        if ( (number & WINK) == WINK){
            handShakeList.add(Signal.WINK);
        }

        if ( (number & DOUBLE_BLINK) == DOUBLE_BLINK){ 
            handShakeList.add(Signal.DOUBLE_BLINK); 
        }

        if ( (number & CLOSE_YOUR_EYES) == CLOSE_YOUR_EYES){
            handShakeList.add(Signal.CLOSE_YOUR_EYES);
            
        }

        if ( (number & JUMP) == JUMP){
            handShakeList.add(Signal.JUMP); 
        }

        if ( (number & REVERSE) == REVERSE){
            
            reverseList = true;
        }

        if (reverseList){
            Collections.reverse(handShakeList);
        }
        return handShakeList;
    }
}
