package grid;

public class Distance {

    public static int findManhattanDistance(Coords c1, Coords c2) {
        int width = Math.abs(c1.x - c2.x);
        int height = Math.abs(c1.y - c2.y);

        return width + height;
    }

    public static int findWeight(Coords c1, Coords c2) {
        return c1.steps + c2.steps;
    }
}
