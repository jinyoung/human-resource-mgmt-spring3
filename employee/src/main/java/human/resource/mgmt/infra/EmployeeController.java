package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import java.util.Optional;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
// @RequestMapping(value="/employees")
@Transactional
public class EmployeeController {

    @Autowired
    EmployeeRepository employeeRepository;

    @RequestMapping(
        value = "employees/{id}/resign",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public Employee resign(
        @PathVariable(value = "id") String id,
        @RequestBody ResignCommand resignCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /employee/resign  called #####");
        Optional<Employee> optionalEmployee = employeeRepository.findById(id);

        optionalEmployee.orElseThrow(() -> new Exception("No Entity Found"));
        Employee employee = optionalEmployee.get();
        employee.resign(resignCommand);

        employeeRepository.save(employee);
        return employee;
    }
}
