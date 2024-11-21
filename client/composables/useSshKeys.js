export default function useSshKeys() {
  const sshKeys = useState("sshKeys", () => []);

  return { sshKeys };
}