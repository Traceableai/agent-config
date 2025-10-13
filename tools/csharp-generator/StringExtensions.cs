using System.Text.RegularExpressions;

public static class StringExtensions
{
  public static string ToUpperSnakeCase(this string input)
  {
    if (string.IsNullOrEmpty(input))
      return input;

    // Insert underscore between lowercase/number and uppercase letter
    string result = Regex.Replace(input, @"(?<=[a-z0-9])([A-Z])", "_$1");

    // Insert underscore before transition from acronym to word
    result = Regex.Replace(result, @"(?<=[A-Z])([A-Z][a-z])", "_$1");

    return result.ToUpperInvariant();
  }

  public static string GetTypeName(this string input)
  {
    int lastDot = input.LastIndexOf('.');
    return lastDot >= 0 ? input.Substring(lastDot + 1) : input;
  }
  
public static string ToPascalCase(this string name)
{
  return string.Concat(name.Split('_')
      .Select(s => char.ToUpperInvariant(s[0]) + s.Substring(1)));
}
}
