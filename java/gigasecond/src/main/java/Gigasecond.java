import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

class Gigasecond {

    // A gigasecond is 10^9 (1,000,000,000) seconds.
    private static final double GIGASECONDS = 1e9;
    private LocalDateTime birthDateTime;

    Gigasecond(LocalDate birthDate) {
        this.birthDateTime =  LocalDateTime.of(birthDate,
                                               LocalTime.of(0,0,0,0)
                                               );
    }

    Gigasecond(LocalDateTime birthDateTime) {
        this.birthDateTime = birthDateTime;
    }

    LocalDateTime getDate() {
        return this.birthDateTime.plusSeconds((long) Gigasecond.GIGASECONDS);
    }

}
