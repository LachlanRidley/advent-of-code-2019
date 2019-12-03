package grid;

public class Instruction {
    public final char dir;
    public final int length;

    public Instruction(char dir, int length) {
        this.dir = dir;
        this.length = length;
    }

    public static Instruction from(String instruction) {
        return new Instruction(instruction.charAt(0), Integer.parseInt(instruction.substring(1)));
    }
}
