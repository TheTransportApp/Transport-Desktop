import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-server-manager',
  templateUrl: './server-manager.component.html'
})
export class ServerManagerComponent implements OnInit {

  public modalOpened: boolean = false;
  public editingServer: boolean = false;

  public editingServerName: string;

  public deleteConfirmation: boolean = false;
  public deletingServer: string = "";

  public servers = [
    {
      name: "Mein Server ",
      address: "192.168.2.1",
      added: "10.06.2018",
      lastconnected: "11.06.2018",
      auth: "Kennwort",
      username: "root",
      password: "test"
    },
  ];

  public formName: string;
  public formAddress: string;
  public formUsername: string;
  public formPassword: string;

  constructor() { }

  ngOnInit() {
  }

  openModal() {
    this.modalOpened = true;
  }

  closeModal() {
    this.formName = this.formAddress = this.formUsername = this.formPassword = "";
    this.modalOpened = false;
  }

  addServer() {
    if(this.formName === "") return;
    if(this.formAddress === "") return;
    if(this.formUsername === "") return;
    if(this.formPassword === "") return;

    this.servers.push({
      name: this.formName,
      address: this.formAddress,
      added: "17.06.2018",
      lastconnected: "-",
      auth: "Kennwort",
      username: this.formUsername,
      password: this.formPassword
    });

    this.closeModal();
  }

  removeServer(name) {
    this.servers = this.servers.filter((item) => item.name !== name);
    this.closeDeleteConfirmation();
  }

  editServer(name) {
    this.editingServer = true;
    this.editingServerName = name;

    let server = this.servers.find(s => s.name === name);

    this.formName = server.name;
    this.formAddress = server.address;
    this.formUsername = server.username;
    this.formPassword = server.password;
  }

  closeEditModal() {
    this.editingServerName = "";
    this.formName = this.formAddress = this.formUsername = this.formPassword = "";
    this.editingServer = false;
  }

  saveEdit() {
    if(this.formName === "") return;
    if(this.formAddress === "") return;
    if(this.formUsername === "") return;
    if(this.formPassword === "") return;

    this.servers = this.servers.map(s => {
      if(s.name === this.editingServerName) {
        s.name = this.formName;
        s.address = this.formAddress;
        s.username = this.formUsername;
        s.password = this.formPassword;
      }
      return s;
    });
    this.closeEditModal();
  }

  openDeleteConfirmation(name) {
    this.deleteConfirmation = true;
    this.deletingServer = name;
  }

  closeDeleteConfirmation() {
    this.deleteConfirmation = false;
    this.deletingServer = "";
  }

}
