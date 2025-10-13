using System.Text;
using Google.Protobuf;
using Google.Protobuf.Compiler;
using Google.Protobuf.Reflection;

var req = CodeGeneratorRequest.Parser.ParseFrom(Console.OpenStandardInput());
var res = new CodeGeneratorResponse();

foreach (var file in req.ProtoFile)
{
  if (file.Name != "proto/ai/traceable/agent/config/v1/config.proto") continue;
  foreach (var msg in file.MessageType)
  {
    var sb = new StringBuilder();
    string className = msg.Name;
    sb.AppendLine("using System;");
    sb.AppendLine("using Tracaeble.Agent.Config.V1;");
    sb.AppendLine("namespace Tracaeble.Agent.Config.Loader;");
    sb.AppendLine($"internal class {className}EnvLoader");
    sb.AppendLine("{");
    sb.AppendLine("    private readonly string envVarPrefix;");
    sb.AppendLine($"    internal {className}EnvLoader (string envVarPrefix)");
    sb.AppendLine("    {");
    sb.AppendLine($"        this.envVarPrefix = envVarPrefix;");
    sb.AppendLine("    }");
    sb.AppendLine($"    internal {className} FromEnv()");
    sb.AppendLine("     {");
    sb.AppendLine($"        var cfg = new {className}();");

    foreach (var field in msg.Field)
    {
      string genName = field.Name.ToPascalCase();
      string constructor = field.Type switch
      {
        FieldDescriptorProto.Types.Type.Enum => $"Enum.TryParse(Environment.GetEnvironmentVairable($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\"), out var {genName.ToLower()}) ? {genName.ToLower()} : {field.TypeName.GetTypeName()}.Unspecified",
        FieldDescriptorProto.Types.Type.Message =>
          field.TypeName switch
          {
            ".google.protobuf.StringValue" => $"Environment.GetEnvironmentVariable($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\")",
            ".google.protobuf.Int32Value" => $"Int32.TryParse(Environment.GetEnvironmentVariable($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\"), out var {genName.ToLower()})? {genName.ToLower()} : null",
            ".google.protobuf.Int64Value" => $"Int64.TryParse(Environment.GetEnvironmentVariable($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\"), out var {genName.ToLower()})? {genName.ToLower()} : null",
            ".google.protobuf.BoolValue" => $"bool.TryParse(Environment.GetEnvironmentVariable($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\"), out var {genName.ToLower()}) ? {genName.ToLower()} : null",
            _ => $"new {genName}EnvLoader($\"{{envVarPrefix}}_{field.Name.ToUpperSnakeCase()}\").FromEnv()",
          },
        _ => throw new ArgumentException($"Unsupported field type {field.Type}"),
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

