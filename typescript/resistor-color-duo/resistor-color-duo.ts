export class ResistorColor {
  private colors: string[];
  private values: Map<string, string>;

  constructor(colors: string[]) {
    if (colors.length >=2){
      this.colors = colors;
      this.values = new Map<string, string>();
      this.values.set("black", "0")
      this.values.set("brown", "1")
      this.values.set("orange", "3")
      this.values.set("yellow", "4")
      this.values.set("green", "5")
      this.values.set("blue", "6")
      this.values.set("violet", "7")
      this.values.set("grey", "8")
    } else {
      throw new Error("At least two colors need to be present")
    }
  }
  
  value = (): number => {
    let result = ""
    for( const color of this.colors){
      const value  = this.values.get(color)
      result += value
      if (result.length == 2) {
        break
      }
    }
    return Number(result);
  }
}
