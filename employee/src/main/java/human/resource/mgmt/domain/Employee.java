package human.resource.mgmt.domain;

import human.resource.mgmt.EmployeeApplication;
import human.resource.mgmt.domain.EmployeeJoined;
import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;

@Entity
@Table(name = "Employee_table")
@Data
public class Employee {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private String userId;

    private String name;

    private String email;

    @PostPersist
    public void onPostPersist() {
        EmployeeJoined employeeJoined = new EmployeeJoined(this);
        employeeJoined.publishAfterCommit();
    }

    public static EmployeeRepository repository() {
        EmployeeRepository employeeRepository = EmployeeApplication.applicationContext.getBean(
            EmployeeRepository.class
        );
        return employeeRepository;
    }

    public void resign(ResignCommand resignCommand) {
        EmployeeResigned employeeResigned = new EmployeeResigned(this);
        employeeResigned.publishAfterCommit();
    }
}
