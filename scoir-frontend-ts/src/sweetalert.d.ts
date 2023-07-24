// sweetalert.d.ts

declare module 'sweetalert' {
    function swal(
      title: string,
      message?: string,
      icon?: string
    ): Promise<any>;
  
    export = swal;
  }
  