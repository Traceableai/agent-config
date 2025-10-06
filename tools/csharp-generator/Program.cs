using System.Text;
using Google.Protobuf;
using Google.Protobuf.Compiler;
using Google.Protobuf.Reflection;

string EnvVarPrefix = "TA";
var req = CodeGeneratorRequest.Parser.ParseFrom(Console.OpenStandardInput());
var res = new CodeGeneratorResponse();

foreach (var file in req.ProtoFile)
{
  foreach (var msg in file.MessageType)
  {
    var sb = new StringBuilder();
    string className = msg.Name;
    sb.AppendLine("using System;");
    sb.AppendLine($"public static class {className}EnvLoader");
    sb.AppendLine("{");
    sb.AppendLine($"    public static {className} FromEnv()");
    sb.AppendLine("     {");
    sb.AppendLine($"        var cfg = new {className}();");

    foreach (var field in msg.Field)
    {
      string envVar = $"{EnvVarPrefix}_{className.ToUpperInvariant()}_{field.Name.ToUpperInvariant()}";
      string genName = ToPascalCase(field.Name);
      string constructor = field.TypeName switch
      {
        ".google.protobuf.StringValue" => $"Environment.GetEnvironmentVariable(\"{envVar}\")",
        ".google.protobuf.Int32Value" => $"Int32.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {genName.ToLower()})? {genName.ToLower()} : null",
        ".google.protobuf.Int64Value" => $"Int64.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {genName.ToLower()})? {genName.ToLower()} : null",
        ".google.protobuf.BoolValue" => $"bool.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {genName.ToLower()}) ? {genName.ToLower()} : null",
        _ => ""
      };
      sb.AppendLine($"        cfg.{genName} = {constructor};");
    }

    sb.AppendLine($"        return cfg;");
    sb.AppendLine("     }");
    sb.AppendLine("}");
    res.File.Add(new CodeGeneratorResponse.Types.File
    {
      Name = $"{className}EnvLoader.cs",
      Content = sb.ToString()
    });
  }
}

res.WriteTo(Console.OpenStandardOutput());

string ToPascalCase(string name)
{
  return string.Concat(name.Split('_')
      .Select(s => char.ToUpperInvariant(s[0]) + s.Substring(1)));
}

string GenerateConstructorRecursively(string fieldPrefix, string envVarPrefix, FieldDescriptorProto field, StringBuilder sb)
{
  string envVar = $"{envVarPrefix}_{field.Name.ToUpperInvariant()}";
  string genName = ToPascalCase(field.Name);
  string fieldKey = $"{fieldPrefix}.{genName}";
  return field.TypeName switch
  {
    ".google.protobuf.StringValue" => $"Environment.GetEnvironmentVariable(\"{envVar}\")",
    ".google.protobuf.Int32Value" => $"Int32.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {field.Name.ToLower()})? {field.Name.ToLower()} : null",
    ".google.protobuf.Int64Value" => $"Int64.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {field.Name.ToLower()})? {field.Name.ToLower()} : null",
    ".google.protobuf.BoolValue" => $"bool.TryParse(Environment.GetEnvironmentVariable(\"{envVar}\"), out var {field.Name.ToLower()}) ? {field.Name.ToLower()} : null",
    _ => ""
  };
}

