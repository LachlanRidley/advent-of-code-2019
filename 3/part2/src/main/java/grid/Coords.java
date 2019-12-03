package grid;

import java.util.Objects;

public class Coords {
    public final int x;
    public final int y;
    public final int steps;

    public Coords(int x, int y, int steps) {
        this.x = x;
        this.y = y;
        this.steps = steps;
    }

    public static final Coords ORIGIN = new Coords(0, 0, 0);

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Coords coords = (Coords) o;
        return x == coords.x &&
                y == coords.y;
    }

    @Override
    public int hashCode() {
        return Objects.hash(x, y);
    }

    @Override
    public String toString() {
        return "Coords{" +
                "x=" + x +
                ", y=" + y +
                ", steps=" + steps +
                '}';
    }
}
