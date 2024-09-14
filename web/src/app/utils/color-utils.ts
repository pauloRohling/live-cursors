export type RGB = { red: number; green: number; blue: number };

export class ColorUtils {
  static getRGB(color: string): RGB {
    const hex = color.replace("#", "");

    return {
      red: Number.parseInt(hex.substring(0, 2), 16),
      green: Number.parseInt(hex.substring(2, 4), 16),
      blue: Number.parseInt(hex.substring(4, 6), 16),
    };
  }

  static getBrightness(color: string): number {
    const { red, green, blue } = ColorUtils.getRGB(color);
    return (red * 299.0 + green * 587.0 + blue * 114.0) / 1000.0;
  }

  static isLight(color: string): boolean {
    const brightness = ColorUtils.getBrightness(color);
    return brightness > 155;
  }
}
