export class ResistorColor {
  private colors: string[];
  private resistorsValues: Map<string, string>;

  constructor(colors: string[]) {
    if (this.isValid(colors)) {
      this.colors = colors;
      this.resistorsValues = this.loadResistorValues();
    } else {
      throw new Error("At least two colors need to be present");
    }
  }

  value = (): number => {
    let result = "";
    for (const color of this.colors) {
      result += this.resistorsValues.get(color);
      if (result.length == 2) {
        break;
      }
    }
    return Number(result);
  };

  private isValid = (colors: string[]): boolean => colors.length >= 2;

  private loadResistorValues = (): Map<string, string> => {
    const resitorsValues = new Map<string, string>();
    resitorsValues.set("black", "0");
    resitorsValues.set("brown", "1");
    resitorsValues.set("orange", "3");
    resitorsValues.set("yellow", "4");
    resitorsValues.set("green", "5");
    resitorsValues.set("blue", "6");
    resitorsValues.set("violet", "7");
    resitorsValues.set("grey", "8");
    return resitorsValues;
  };
}
