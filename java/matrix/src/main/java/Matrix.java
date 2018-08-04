
class Matrix {

    private int [ ] [ ] data;
    private int rowsCount;
    private int columnsCount;
    private static final String LINE_SEPARATOR = "\n";
    private static final String COLUMN_SEPARATOR = " ";


    Matrix(String matrixAsString) {
        String [ ] rowsArray = matrixAsString.split(LINE_SEPARATOR);
        this.rowsCount = rowsArray.length;
        String [ ] columnsArray;

        for (int i = 0; i < this.getRowsCount(); i++) {

            columnsArray = rowsArray[i].split(COLUMN_SEPARATOR);
            this.columnsCount = columnsArray.length;
            if (this.data == null) {
                this.data = new int[this.getRowsCount()][this.getColumnsCount()];                
            }
            for (int j = 0; j < this.getColumnsCount(); j++) {
                this.data[i][j] = Integer.parseInt(columnsArray[j]);
            }
            
        }

    }

    int[] getRow(int rowNumber) {

        return this.data[rowNumber];

    }

    int[] getColumn(int columnNumber) {

        int[] columnsArray = new int[this.getRowsCount()];
        for (int i = 0; i < columnsArray.length; i++) {
            columnsArray[i] = this.data[i][columnNumber];
        }
        return columnsArray;
    }

    int getRowsCount() {
        return this.rowsCount;
    }

    int getColumnsCount() {
        return this.columnsCount;
    }
}
